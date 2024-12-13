import { Component, computed, inject } from '@angular/core';
import {PrimaryButtonComponent} from '../../../components/primary-button/primary-button.component';
import {CartService} from '../../../services/cart.service';

@Component({
  selector: 'app-order-summary',
  standalone: true,
  imports: [PrimaryButtonComponent],
  templateUrl: './order-summary.component.html',
  styles: ``
})
export class OrderSummaryComponent {
  cartService = inject(CartService);

  total = computed(() => {
    let total = 0;
    for (const item of this.cartService.cart()) {
      total += item.price;
    }

    return total;
  });
}
