import {Component, inject} from '@angular/core';
import {NavComponent} from './nav/nav.component';
import {Router, RouterLink} from '@angular/router';
import {CartService} from '../../services/cart.service';
import {PrimaryButtonComponent} from '../primary-button/primary-button.component';
import {NgIf} from '@angular/common';

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [
    RouterLink,
    PrimaryButtonComponent,
    NgIf
  ],
  templateUrl: './header.component.html',
  styles: ``
})
export class HeaderComponent {
  constructor(private router: Router) {
  }
  cartService = inject(CartService);

  isLandingPage(): boolean {
    return this.router.url === '/';
  }

  isProductsPage(): boolean {
    return this.router.url === '/product';
  }

  isCartCheckout(): boolean {
    return this.router.url === '/cart';
  }

  onSignOut() {
    console.log('Signing out...');
    // Add your sign-out logic (e.g., clear tokens, redirect)
    this.router.navigate(['/auth']);
  }
  showBtnClicked() {
    console.log('showBtnClicked');
  }
}
