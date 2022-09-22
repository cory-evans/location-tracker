import { NgModule } from '@angular/core';
import { ListComponent } from './list/list.component';
import { DevicesRoutingModule } from './devices.routes';
import { SharedModule } from '../shared/shared.module';
import { CreateComponent } from './create/create.component';
import { DevicesComponent } from './devices.component';

@NgModule({
  declarations: [ListComponent, CreateComponent, DevicesComponent],
  imports: [SharedModule, DevicesRoutingModule],
})
export class DevicesModule {}
