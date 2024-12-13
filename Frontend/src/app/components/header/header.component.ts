import {Component, inject} from '@angular/core';
import {NavComponent} from './nav/nav.component';
import {RouterLink} from '@angular/router';
import {CartService} from '../../services/cart.service';
import {PrimaryButtonComponent} from '../primary-button/primary-button.component';

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [
    NavComponent,
    RouterLink,
    PrimaryButtonComponent
  ],
  templateUrl: './header.component.html',
  styles: ``
})
export class HeaderComponent {
  cartService = inject(CartService);

  showBtnClicked() {
    console.log('showBtnClicked');
  }
}
