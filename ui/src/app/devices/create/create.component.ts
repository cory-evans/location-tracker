import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PocketbaseService } from 'src/app/shared/services/pocketbase.service';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
})
export class CreateComponent implements OnInit {
  form: FormGroup;
  constructor(fb: FormBuilder, private api: PocketbaseService) {
    this.form = fb.group({
      name: [null, Validators.required],
    });
  }

  ngOnInit(): void {}

  onSubmit() {
    if (this.form.invalid) {
      return;
    }

    const ownerId = this.api.authStore.model?.id;
    if (!ownerId) {
      return;
    }

    const name = this.form.get('name')?.value;

    const device = {
      owner: ownerId,
      name: name,
    };

    this.api.records
      .create('device', device)
      .then((r) => {
        console.log(r);
      })
      .catch((e) => console.error(e));
  }
}
