import { Injectable } from '@angular/core';
import {
  ActivatedRouteSnapshot,
  CanActivate,
  CanActivateChild,
  CanLoad,
  Route,
  Router,
  RouterStateSnapshot,
  UrlSegment,
} from '@angular/router';
import { PocketbaseService } from '../services/pocketbase.service';

@Injectable({
  providedIn: 'root',
})
export class AuthenticatedGuard
  implements CanActivate, CanActivateChild, CanLoad
{
  constructor(private api: PocketbaseService, private router: Router) {}

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
    return this.isAuthenticated();
  }

  canActivateChild(
    childRoute: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ) {
    return this.isAuthenticated();
  }

  canLoad(route: Route, segments: UrlSegment[]) {
    return this.isAuthenticated();
  }

  isAuthenticated() {
    if (!this.api.authStore.isValid) {
      return this.router.createUrlTree(['auth', 'signin']);
    }

    return true;
  }
}
