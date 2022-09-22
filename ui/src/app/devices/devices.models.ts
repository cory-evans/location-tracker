import { Record } from 'pocketbase';

export interface Device extends Record {
  name: string;
}
