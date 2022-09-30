import { Record } from 'pocketbase';

export interface Device extends Record {
  name: string;
}

export interface Location extends Record {
  lat: number;
  lon: number;
}

export interface LocationWithDevice extends Location {
  '@expand': {
    device: Device;
  };
}
