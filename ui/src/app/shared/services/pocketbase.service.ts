import { Injectable } from '@angular/core';
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

  get locations() {
    return this.client.collection('locations');
  }

  get devices() {
    return this.client.collection('device');
  }

  get users() {
    return this.client.collection('users');
  }

  get authStore() {
    return this.client.authStore;
  }

  get myid() {
    return this.client.authStore.model?.id || ''
  }

  createDeviceToken(deviceId: string) {
    return this.client.send(`/api/device/${deviceId}/token`, {
      method: 'GET',
    });
  }
}
