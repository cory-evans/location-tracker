import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { PocketbaseService } from 'src/app/shared/services/pocketbase.service';

@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  host: {
    class: 'flex-1',
  },
})
export class SignInComponent implements OnInit {
  form: FormGroup;
  validationError?: string;
  constructor(
    fb: FormBuilder,
    private api: PocketbaseService,
    private router: Router
  ) {
    this.form = fb.group({
      email: [null, Validators.required],
      password: [null, Validators.required],
    });
  }

  ngOnInit(): void {
    this.navigateAwayIfSignedIn();
  }

  navigateAwayIfSignedIn() {
    if (!this.api.authStore.isValid) {
      return;
    }

    this.router.navigate(['map']);
  }

  onSubmit() {
    if (this.form.invalid) {
      Object.entries(this.form.controls).forEach(([key, control]) => {
        control.markAsDirty();
      });

      return;
    }

    const email = this.form.get('email')!.value;
    const password = this.form.get('password')!.value;

    this.api.users
      .authWithPassword(email, password)
      .then(() => {
        this.navigateAwayIfSignedIn();
      })
      .catch((error) => {
        this.validationError = error.message;
      });
  }
}
