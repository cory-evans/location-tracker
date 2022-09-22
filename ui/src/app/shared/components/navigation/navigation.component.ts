import { Component, OnInit } from '@angular/core';
import {
  faSnowflake,
  faMap,
  faUserCircle,
} from '@fortawesome/free-regular-svg-icons';
import { faMicrochip } from '@fortawesome/free-solid-svg-icons';
import type { IconDefinition } from '@fortawesome/fontawesome-common-types';

import { NgxPopperjsPlacements, NgxPopperjsTriggers } from 'ngx-popperjs';
import { PocketbaseService } from '../../services/pocketbase.service';

interface MenuItem {
  link: string | string[];
  icon: IconDefinition;
}

@Component({
  selector: 'app-navigation',
  templateUrl: './navigation.component.html',
})
export class NavigationComponent implements OnInit {
  brandIcon = faSnowflake;
  accountIcon = faUserCircle;
  placements = NgxPopperjsPlacements;
  triggers = NgxPopperjsTriggers;

  constructor(private api: PocketbaseService) {}

  get showNav() {
    return this.api.authStore.isValid;
  }

  menuItems: MenuItem[] = [
    {
      icon: faMap,
      link: '/map',
    },
    {
      icon: faMicrochip,
      link: '/devices',
    },
  ];

  ngOnInit(): void {}

  userSignOut() {
    this.api.authStore.clear();
  }
}
