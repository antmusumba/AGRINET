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
  styles: ``
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
