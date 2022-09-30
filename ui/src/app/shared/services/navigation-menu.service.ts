import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class NavigationMenuService {
  constructor() {}

  public isOpen = false;

  openMenu() {
    this.isOpen = true;
  }

  closeMenu() {
    this.isOpen = false;
  }
}
