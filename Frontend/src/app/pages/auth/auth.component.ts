import {AbstractControl, FormBuilder, FormGroup, ReactiveFormsModule, ValidatorFn, Validators} from '@angular/forms';
import {Component} from '@angular/core';
import {CommonModule} from '@angular/common';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import {Router} from '@angular/router';

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
  styles: [`
    .auth-container {
      display: flex;
      justify-content: center;
      align-items: center;
      height: 100vh;
    }

    .auth-card {
      background: rgba(255, 255, 255, 0.85);
      padding: 2rem;
      border-radius: 8px;
      box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
      text-align: center;
      width: 100%;
      max-width: 400px;
    }

    h1 {
      font-size: 1.8rem;
      color: #333;
      margin-bottom: 1rem;
    }

    .highlight {
      color: #ff9800;
    }

    .link {
      color: #007bff;
      text-decoration: underline;
      cursor: pointer;
    }

    form {
      display: flex;
      flex-direction: column;
      gap: 1rem;
    }

    .form-group {
      display: flex;
      flex-direction: column;
      align-items: flex-start;
    }

    .form-group label {
      font-size: 0.9rem;
      margin-bottom: 0.5rem;
    }

    .form-group input {
      width: 100%;
      padding: 0.5rem;
      border: 1px solid #ccc;
      border-radius: 4px;
      font-size: 1rem;
    }

    .form-group input.invalid {
      border-color: #e74c3c;
    }

    .form-group small {
      color: #e74c3c;
      font-size: 0.8rem;
    }

    .form-group-row {
      display: flex;
      gap: 1rem;
    }

    .btn-submit {
      background-color: #4caf50;
      color: #fff;
      border: none;
      padding: 0.8rem 1rem;
      font-size: 1rem;
      border-radius: 4px;
      cursor: pointer;
    }

    .btn-submit:hover {
      background-color: #43a047;
    }


  `]
})
export class AuthComponent {
  authForm: FormGroup;
  isSignUp = false;

  constructor(private fb: FormBuilder, private router: Router) {
    this.authForm = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]],
      username: [''], // Only for Sign Up
      firstName: [''], // Only for Sign Up
      lastName: [''], // Only for Sign Up
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
      const formData = { ...this.authForm.value };
      if (!this.isSignUp) {
        delete formData.username;
        delete formData.firstName;
        delete formData.lastName;
      }
      console.log('Form Data:', formData);

      // Navigate to the `/product` page
      this.router.navigate(['/product']);
    } else {
      console.log('Form is invalid!');
    }
  }
}
