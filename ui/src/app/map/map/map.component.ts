import { Component, OnInit } from '@angular/core';

import { tileLayer, latLng, marker, Marker, icon, Map } from 'leaflet';
import { Device, Location } from 'src/app/shared/models';
import { PocketbaseService } from 'src/app/shared/services/pocketbase.service';

@Component({
  selector: 'app-map',
  templateUrl: './map.component.html',
  host: {
    class: 'flex-1',
  },
})
export class MapComponent implements OnInit {
  constructor(private api: PocketbaseService) {}
  map?: Map;
  devices: Device[] = [];
  locations: {
    [deviceId: string]: Location[];
  } = {};

  bottomMenuHidden = true;

  ngOnInit(): void {
    this.api.devices
      .getFullList()
      .then((devices) => {
        return devices as Device[];
      })
      .then(async (devices) => {
        this.devices = devices;
        for (let i = 0; i < devices.length; i++) {
          const device = devices[i];
          await this.api.locations.getList(1, 5, {
              sort: '-created',
              filter: `device = '${device.id}'`,
            })
            .then((data) => {
              return data.items as Location[];
            })
            .then((items) => {
              return items.map((i) => {
                i.created = i.created + 'Z';
                return i;
              });
            })
            .then((items) => {
              this.locations[device.id] = items;
              items.forEach((item) => {
                const l = item as Location;
                this.markers.push(
                  marker([l.lat, l.lon], {
                    icon: icon({
                      iconSize: [25, 41],
                      iconAnchor: [13, 41],
                      iconUrl: 'assets/marker-icon.png',
                      iconRetinaUrl: 'assets/marker-icon-2x.png',
                      shadowUrl: 'assets/marker-shadow.png',
                    }),
                  })
                );
              });
            });
        }
      });
  }

  getLatestLocation(d: Device) {
    if (!this.locations[d.id] || !this.locations[d.id].length) return null;

    return this.locations[d.id][0];
  }

  markers: Marker[] = [];

  options = {
    layers: [
      tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 18,
        attribution: '...',
      }),
    ],
    zoom: 10,
    center: latLng(-39.53688065064306, 176.7459869384766),
  };

  onMapReady(map: Map) {
    this.map = map;
    map.zoomControl.setPosition('topright');
  }

  goto(d: Device) {
    if (this.locations[d.id] === undefined) {
      return;
    }

    const l = this.locations[d.id][0];
    this.map?.flyTo([l.lat, l.lon]);

    this.bottomMenuHidden = true;
  }

  toggleBottomMenu() {
    this.bottomMenuHidden = !this.bottomMenuHidden;
  }
}
