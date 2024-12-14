import { Component, signal } from '@angular/core';
import { Product } from '../../models/products.models';
import { ProductCardComponent } from './product-card/product-card.component';
import { SearchBarComponent } from '../../components/search-bar/search-bar.component';
import { FilterProductsPipe } from '../../pipes/filter-products.pipe';
import {NgIf} from '@angular/common';

@Component({
  selector: 'app-products-list',
  standalone: true,
  imports: [ProductCardComponent, SearchBarComponent, FilterProductsPipe, NgIf],
  templateUrl: './products-list.component.html',
  styles: [],
})
export class ProductsListComponent {
  products = signal<Product[]>([
    {
      id: 1,
      title: "Farm Tractor",
      price: 109.95,
      image: "../hero-image.png",
      stock: 0,
    },
    {
      id: 2,
      title: "Shovel",
      price: 39.99,
      image: "",
      stock: 10,
    },
    {
      id: 3,
      title: "Irrigation System",
      price: 999.99,
      image: "https://image.tmdb.org/t/p/w500",
      stock: 10,
    },
    {
      id: 4,
      title: "Organic Seeds",
      price: 19.95,
      image: "https://image.tmdb.org/t/p/w500",
      stock: 0,
    },
  ]);

  searchQuery = signal<string>(''); // Holds the search query

  onSearch(query: string) {
    this.searchQuery.set(query); // Update the search query when the user types
  }
}
