import { Component, OnInit } from '@angular/core';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { ProductsService } from '../../../services/products.service';
import { Product } from '../../../models/products.models';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-product-form',
  templateUrl: './product-form.component.html',
  styleUrls: ['./product-form.component.css'],
  imports: [
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatCardModule,
    ReactiveFormsModule,
    CommonModule,
  ],
})
export class AddProductComponent implements OnInit {
  addProductForm!: FormGroup;

  constructor(
    private fb: FormBuilder,
    private productService: ProductsService,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.addProductForm = this.fb.group({
      title: ['', [Validators.required]],
      description: ['', [Validators.required, Validators.minLength(10)]],
      price: ['', [Validators.required, Validators.min(0)]],
      stock: ['', [Validators.required, Validators.min(0)]],
      image: [''],
    });
  }

  onAddProduct(): void {
    if (this.addProductForm.valid) {
      const newProduct: Product = this.addProductForm.value;
      this.productService.createProduct(newProduct).subscribe({
        next: (addedProduct) => {
          console.log('Product added successfully:', addedProduct);
          this.router.navigate(['/products']);
        },
        error: (err) => {
          console.error('Error adding product:', err);
        },
      });
    }
  }
}
