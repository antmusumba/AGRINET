import { Component, computed, inject } from '@angular/core';
import { PrimaryButtonComponent } from '../../../components/primary-button/primary-button.component';
import { CartService } from '../../../services/cart.service';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-order-summary',
  standalone: true,
  imports: [PrimaryButtonComponent, CommonModule],
  templateUrl: './order-summary.component.html',
})
export class OrderSummaryComponent {
  cartService = inject(CartService);
  isLoading = false;
  errorMessage: string | null = null;

  constructor(private router: Router) {}

  total = computed(() => {
    let total = 0;
    for (const item of this.cartService.cart()) {
      total += item.price;
    }
    return total;
  });

  handleCheckout() {
    this.isLoading = true;
    this.errorMessage = null;

    this.cartService.initiatePayment('254720804060', this.total()).subscribe({
      next: (response) => {
        this.isLoading = false;
        console.log('Payment Success:', response);
        this.router.navigate(['/product']);
      },
      error: (err) => {
        this.isLoading = false;
        this.errorMessage = err.message || 'Oops! Something went wrong.';
        console.error('Payment Error:', err);
      },
    });
  }
}
