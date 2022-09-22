import { Component, Input, OnInit } from '@angular/core';

export interface MenuItem {
  path: string | any[];
  text: string;
}
export type MenuItems = MenuItem[];

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
})
export class SidebarComponent implements OnInit {
  @Input() title?: string;
  @Input() items: MenuItem[] = [];
  constructor() {}

  ngOnInit(): void {}
}
