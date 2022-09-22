import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MapComponent } from './map/map.component';
import { SharedModule } from '../shared/shared.module';
import { MapRoutingModule } from './map.routes';

import { LeafletModule } from '@asymmetrik/ngx-leaflet';

@NgModule({
  declarations: [MapComponent],
  imports: [CommonModule, SharedModule, MapRoutingModule, LeafletModule],
})
export class MapModule {}
