import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import Pocketbase from 'pocketbase';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root',
})
export class PocketbaseService {
  private client: Pocketbase;
  constructor() {
    this.client = new Pocketbase(environment.pocketbase);
  }

  get records() {
    return this.client.records;
  }

  get users() {
    return this.client.users;
  }

  get authStore() {
    return this.client.authStore;
  }

  createDeviceToken(deviceId: string) {
    return this.client.send(`/api/device/${deviceId}/token`, {
      method: 'GET',
    });
  }
}
