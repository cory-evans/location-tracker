import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AccountComponent } from './account/account.component';
import { SettingsRoutingModule } from './settings.routes';

@NgModule({
  declarations: [AccountComponent],
  imports: [CommonModule, SettingsRoutingModule],
})
export class SettingsModule {}
