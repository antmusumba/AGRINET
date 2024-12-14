import { Component, OnInit, signal } from '@angular/core';
import { Product } from '../../models/products.models';
import { ProductCardComponent } from './product-card/product-card.component';
import { SearchBarComponent } from '../../components/search-bar/search-bar.component';
import { FilterProductsPipe } from '../../pipes/filter-products.pipe';
import { NgIf } from '@angular/common';
import { ProductsService } from '../../services/products.service';

@Component({
  selector: 'app-products-list',
  standalone: true,
  imports: [ProductCardComponent, SearchBarComponent, FilterProductsPipe, NgIf],
  templateUrl: './products-list.component.html',
})
export class ProductsListComponent implements OnInit {
  // Declare the signal for search query and products
  searchQuery = signal<string>('');
  products = signal<Product[]>([]);
  filteredProducts = signal<Product[]>([]);

  constructor(private productsService: ProductsService) {}

  ngOnInit() {
    this.productsService.getAllProducts().subscribe({
      next: (products) => {
        console.log('Fetched products:', products);
        this.products.set(products);
        this.filteredProducts.set(products); // Initially display all products
      },
      error: (err) => console.error('Error fetching products:', err),
    });
  }

  onSearch(query: string) {
    // Update the search query signal and trigger filtering
    this.searchQuery.set(query);
    this.filteredProducts.set(
      this.products().filter((product) =>
        product.title.toLowerCase().includes(query.toLowerCase())
      )
    );
  }
}
