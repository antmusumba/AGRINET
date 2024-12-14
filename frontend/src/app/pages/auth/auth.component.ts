import {
  AbstractControl,
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { Router } from '@angular/router';
import { AuthService } from '../../services/auth.service';

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
})
export class AuthComponent {
  authForm: FormGroup;
  isSignUp = false;
  isLoading = false;
  errorMessage: string | null = null;

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private authService: AuthService
  ) {
    this.authForm = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]],
      firstName: [''],
      lastName: [''],
      phone: ['', [Validators.required]],
      userName: [''],
    });
  }

  toggleMode() {
    this.isSignUp = !this.isSignUp;

    if (!this.isSignUp) {
      this.authForm.patchValue({
        firstName: '',
        lastName: '',
      });
    }
  }

  onSubmit() {
    if (this.authForm.valid) {
      const formData = this.authForm.value;
      this.isLoading = true;
      this.errorMessage = null;

      if (this.isSignUp) {
        this.register(formData);
      } else {
        this.login(formData);
      }
    } else {
      console.log('Form is invalid!');
    }
  }

  private login(formData: any) {
    this.authService.login(formData.email, formData.password).subscribe(
      (response) => {
        this.isLoading = false;
        console.log('Login Success:', response);
        this.router.navigate(['/product']);
      },
      (error) => {
        this.isLoading = false;
        this.errorMessage = error.message || 'Login failed. Please try again.';
        console.error('Login Error:', error);
      }
    );
  }

  private register(formData: any) {
    const registerData = {
      firstName: formData.firstName,
      lastName: formData.lastName,
      email: formData.email,
      password: formData.password,
      phone: formData.phone,
    };

    this.authService.register(registerData).subscribe(
      () => {
        this.isLoading = false;
        console.log('Registration Success');
        this.router.navigate(['/product']);
      },
      (error) => {
        this.isLoading = false;
        this.errorMessage =
          error.message || 'Registration failed. Please try again.';
        console.error('Registration Error:', error);
      }
    );
  }
}
