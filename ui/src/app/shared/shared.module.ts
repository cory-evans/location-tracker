import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';

import { NgxPopperjsModule } from 'ngx-popperjs';

import { NavigationComponent } from './components/navigation/navigation.component';
import { ButtonComponent } from './components/button/button.component';
import { SidebarComponent } from './components/sidebar/sidebar.component';

@NgModule({
  declarations: [NavigationComponent, ButtonComponent, SidebarComponent],
  imports: [CommonModule, ReactiveFormsModule, RouterModule, NgxPopperjsModule],
  exports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    NavigationComponent,
    SidebarComponent,
    ButtonComponent,
  ],
})
export class SharedModule {}
