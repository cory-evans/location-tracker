import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PocketbaseService } from '../shared/services/pocketbase.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  host: {
    class: 'flex-1',
  },
})
export class HomeComponent implements OnInit {
  constructor(private router: Router, private api: PocketbaseService) {
    if (api.myid != '') {
      this.router.navigate(['map'])
    }
  }
  ngOnInit(): void {}

  signIn() {
    this.router.navigate(['auth', 'signin']);
  }
}
