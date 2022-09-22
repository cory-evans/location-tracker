import { Component, OnInit } from '@angular/core';
import { MenuItems } from '../shared/components/sidebar/sidebar.component';

@Component({
  selector: 'app-devices',
  templateUrl: './devices.component.html',
  host: {
    class: 'flex-1 flex flex-col',
  },
})
export class DevicesComponent implements OnInit {
  menuItems: MenuItems = [
    {
      path: 'list',
      text: 'Manage Devices',
    },
    {
      path: 'create',
      text: 'Create Device',
    },
  ];
  constructor() {}

  ngOnInit(): void {}
}
