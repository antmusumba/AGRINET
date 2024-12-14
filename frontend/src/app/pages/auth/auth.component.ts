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
      username: [''],
      firstName: [''],
      lastName: [''],
    });
  }

  toggleMode() {
    this.isSignUp = !this.isSignUp;

    if (!this.isSignUp) {
      this.authForm.patchValue({
        username: '',
        firstName: '',
        lastName: '',
      });
    }
  }

  onSubmit() {
    if (this.authForm.valid) {
      this.login(this.authForm.value);
    } else {
      console.log('Form is invalid!');
    }
  }

  login(formData: any) {
    this.authService.login(formData.email, formData.password).subscribe(
      (response) => {
        this.isLoading = false;
        console.log('Login Success:', response);

        localStorage.setItem('auth_token', response.token);
        localStorage.setItem('user', JSON.stringify(response.user));

        this.router.navigate(['/product']);
      },
      (error) => {
        this.isLoading = false;
        this.errorMessage = 'Login failed. Please check your credentials.';
        console.error('Login Error:', error);
      }
    );
  }

  register(formData: any) {
    const registerData = {
      firstName: formData.firstName,
      lastName: formData.lastName,
      email: formData.email,
      phone: formData.phone,
      password: formData.password,
    };
    this.authService.register(registerData).subscribe(
      () => {
        this.isLoading = false;
        console.log('Registration Success');
        this.router.navigate(['/product']);
      },
      (error) => {
        this.isLoading = false;
        this.errorMessage = 'Registration failed. Please try again.';
        console.error('Registration Error:', error);
      }
    );
  }
}
