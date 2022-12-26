import { Component, OnInit } from '@angular/core';
import { PocketbaseService } from 'src/app/shared/services/pocketbase.service';
import { Device } from 'src/app/shared/models';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  host: {
    class: 'flex-1',
  },
})
export class ListComponent implements OnInit {
  constructor(private api: PocketbaseService) {}

  devices: Device[] = [];

  ngOnInit(): void {
    this.listDevices();
  }

  listDevices() {
    this.api.devices
      .getFullList<Device>(undefined, {
        filter: `owner = "${this.api.myid}"`
      })
      .then((resp) => (this.devices = resp))
  }

  getToken(device: Device) {
    this.api.createDeviceToken(device.id).then((value) => console.log(value));
  }

  deleteDevice(device: Device) {
    this.api.devices.delete(device.id).then(() => {
      this.listDevices()
    })
  }
}
