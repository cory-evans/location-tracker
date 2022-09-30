import { Component, Input, OnInit } from '@angular/core';
import { Map } from 'leaflet';
import { NavigationMenuService } from 'src/app/shared/services/navigation-menu.service';

@Component({
  selector: 'app-control',
  templateUrl: './control.component.html',
})
export class ControlComponent implements OnInit {
  @Input() map?: Map;

  constructor(public navService: NavigationMenuService) {}

  ngOnInit(): void {}
}
