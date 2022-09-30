import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  host: {
    class: 'flex-1',
  },
})
export class HomeComponent implements OnInit {
  constructor(private router: Router) {}
  ngOnInit(): void {}

  signIn() {
    this.router.navigate(['auth', 'signin']);
  }
}
