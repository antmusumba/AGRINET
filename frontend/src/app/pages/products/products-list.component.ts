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
  products = signal<Product[]>([]);
  searchQuery = signal<string>('');

  constructor(private productsService: ProductsService) {}

  ngOnInit() {
    this.productsService.getAllProducts().subscribe({
      next: (products) => this.products.set(products),
      error: (err) => console.error('Error fetching products:', err),
    });
  }

  onSearch(query: string) {
    this.searchQuery.set(query);
  }
}
