import { Component } from '@angular/core';
import {CommonModule} from '@angular/common';
import {FormBuilder, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';

@Component({
  selector: 'app-auth',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatCardModule,
  ],
  templateUrl: './auth.component.html',
  styles: ``
})
export class AuthComponent {
  authForm: FormGroup;
  isSignUp = false;

  constructor(private fb: FormBuilder) {
    this.authForm = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]],
      confirmPassword: [''] // Only for Sign Up
    });
  }

  toggleMode() {
    this.isSignUp = !this.isSignUp;
    if (!this.isSignUp) {
      this.authForm.get('confirmPassword')?.reset();
    }
  }

  onSubmit() {
    if (this.authForm.valid) {
      console.log('Form Data:', this.authForm.value);
    } else {
      console.log('Form is invalid!');
    }
  }

}
