import {Component, inject} from '@angular/core';
import {CartService} from '../../services/cart.service';
import {CartItemComponent} from './cart-item/cart-item.component';


@Component({
  selector: 'app-cart',
  standalone: true,
  imports: [
    CartItemComponent,
  ],
  templateUrl: './cart.component.html',
  styles: ``
})
export class CartComponent {
  cartService = inject(CartService)
}
