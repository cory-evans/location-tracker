import { Component, HostListener } from '@angular/core';

import { PocketbaseService } from '../../services/pocketbase.service';
import { NavigationMenuService } from '../../services/navigation-menu.service';

interface MenuItem {
  link: string | string[];
  icon: string;
  text: string;
}

@Component({
  selector: 'app-navigation',
  templateUrl: './navigation.component.html',
  host: {
    class: 'absolute z-40 left-0 top-0 w-0 h-0',
  },
})
export class NavigationComponent {
  constructor(
    private api: PocketbaseService,
    public navService: NavigationMenuService
  ) {}

  get showNav() {
    return this.api.authStore.isValid;
  }

  clickedInside = false;

  menuItems: MenuItem[] = [
    {
      icon: 'bi bi-map',
      link: '/map',
      text: 'Map',
    },
    {
      icon: 'bi bi-sim',
      link: '/devices',
      text: 'Devices',
    },
  ];

  @HostListener('click')
  clickin() {
    this.clickedInside = true;
  }

  @HostListener('document:click')
  clickout() {
    if (!this.clickedInside) {
      this.navService.closeMenu();
    }
    this.clickedInside = false;
  }

  userSignOut() {
    this.api.authStore.clear();
  }
}
