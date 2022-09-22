import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AccountComponent } from './account/account.component';

const SETTINGS_ROUTES: Routes = [
  {
    path: 'account',
    component: AccountComponent,
  },
];

@NgModule({
  imports: [RouterModule.forChild(SETTINGS_ROUTES)],
  exports: [RouterModule],
})
export class SettingsRoutingModule {}
