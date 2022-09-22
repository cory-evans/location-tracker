import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CreateComponent } from './create/create.component';
import { DevicesComponent } from './devices.component';
import { ListComponent } from './list/list.component';

const DEVICE_ROUTES: Routes = [
  {
    path: '',
    component: DevicesComponent,
    children: [
      {
        path: 'list',
        component: ListComponent,
      },
      {
        path: 'create',
        component: CreateComponent,
      },
      {
        path: '',
        redirectTo: 'list',
        pathMatch: 'full',
      },
    ],
  },
];

@NgModule({
  imports: [RouterModule.forChild(DEVICE_ROUTES)],
  exports: [RouterModule],
})
export class DevicesRoutingModule {}
