import { Component, OnInit } from '@angular/core';
import { Record } from 'pocketbase';
import { PocketbaseService } from '../shared/services/pocketbase.service';
interface Device extends Record {
  name: string;
}
@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss'],
})
export class HomeComponent implements OnInit {
  constructor(private api: PocketbaseService) {}

  devices: Device[] = [];

  ngOnInit(): void {
    this.listDevices();
  }

  listDevices() {
    this.api.records
      .getList('device')
      .then((resp) => (this.devices = resp.items as Device[]));
  }

  getToken(device: Device) {
    this.api.createDeviceToken(device.id).then((value) => console.log(value));
  }
}
