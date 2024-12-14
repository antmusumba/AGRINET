import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject, Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { environment } from '../environment/environment.dev';

export interface User {
  id: string;
  email: string;
  firstName: string;
  lastName: string;
  phone: string;
}

export interface AuthResponse {
  user: User;
  token: string;
}

export interface RegisterRequest {
  firstName: string;
  lastName: string;
  email: string;
  phone: string;
  password: string;
}

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private readonly API_URL = `${environment.apiUrl}/api/auth`;
  private readonly TOKEN_KEY = 'auth_token';
  private readonly USER_KEY = 'user';

  private currentUserSubject = new BehaviorSubject<User | null>(
    this.getStoredUser()
  );
  currentUser$ = this.currentUserSubject.asObservable();

  constructor(private http: HttpClient) {}

  login(email: string, password: string): Observable<AuthResponse> {
    return this.http.post<AuthResponse>(`${this.API_URL}/login`, {
      email,
      password,
    });
  }

  register(userData: RegisterRequest): Observable<void> {
    return this.http
      .post<AuthResponse>(`${this.API_URL}/register`, userData)
      .pipe(map(() => void 0));
  }

  logout(): void {
    localStorage.removeItem(this.TOKEN_KEY);
    localStorage.removeItem(this.USER_KEY);
    this.currentUserSubject.next(null);
  }

  isLoggedIn(): boolean {
    return !!this.getToken();
  }

  getToken(): string | null {
    return localStorage.getItem(this.TOKEN_KEY);
  }

  private getStoredUser(): User | null {
    const user = localStorage.getItem(this.USER_KEY);
    return user ? JSON.parse(user) : null;
  }
}
