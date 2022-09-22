import { Component, OnInit } from '@angular/core';

import { tileLayer, latLng, marker, Marker, icon } from 'leaflet';
import { Record } from 'pocketbase';
import { PocketbaseService } from 'src/app/shared/services/pocketbase.service';

interface LocationRecord extends Record {
  lat: number;
  lon: number;
}

@Component({
  selector: 'app-map',
  templateUrl: './map.component.html',
  host: {
    class: 'flex-1',
  },
})
export class MapComponent implements OnInit {
  constructor(private api: PocketbaseService) {}

  ngOnInit(): void {
    this.api.records.getList('locations').then((data) => {
      data.items.forEach((item) => {
        const l = item as LocationRecord;
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
}
