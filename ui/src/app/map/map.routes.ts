import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { MapComponent } from './map/map.component';

const MAP_ROUTES: Routes = [
  {
    path: '',
    component: MapComponent,
  },
];

@NgModule({
  imports: [RouterModule.forChild(MAP_ROUTES)],
  exports: [RouterModule],
})
export class MapRoutingModule {}
