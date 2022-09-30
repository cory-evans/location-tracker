import { Component, OnInit } from '@angular/core';
import { PocketbaseService } from 'src/app/shared/services/pocketbase.service';
import { Device } from 'src/app/shared/models';

import { faTrashCan } from '@fortawesome/free-regular-svg-icons';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  host: {
    class: 'flex-1',
  },
})
export class ListComponent implements OnInit {
  deleteIcon = faTrashCan;
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

  deleteDevice(device: Device) {
    this.api.records.delete('device', device.id).then((r) => {
      this.listDevices();
    });
  }
}
