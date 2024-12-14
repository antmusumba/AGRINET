import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject, Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { environment } from '../environment/environment.dev';
import { AuthResponse, RegisterRequest, User } from '../models/products.models';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private readonly API_URL = `${environment.apiUrl}/api/auth`;

  private currentUserSubject = new BehaviorSubject<User | null>(
    this.getStoredUser()
  );
  currentUser$ = this.currentUserSubject.asObservable();

  private tokenSubject = new BehaviorSubject<string | null>(
    this.getStoredToken()
  );

  constructor(private http: HttpClient) {}

  login(email: string, password: string): Observable<AuthResponse> {
    return this.http
      .post<AuthResponse>(`${this.API_URL}/login`, { email, password })
      .pipe(
        map((response: AuthResponse) => {
          this.storeAuthData(response.token, response.user);
          return response;
        })
      );
  }

  register(userData: RegisterRequest): Observable<void> {
    return this.http
      .post<AuthResponse>(`${this.API_URL}/register`, userData)
      .pipe(map(() => void 0));
  }

  logout(): void {
    this.clearAuthData();
  }

  isLoggedIn(): boolean {
    return !!this.tokenSubject.value;
  }

  getToken(): string | null {
    return this.tokenSubject.value;
  }

  getUserId(): string | null {
    return this.currentUserSubject.value?.id || null;
  }

  private storeAuthData(token: string, user: User): void {
    // Store auth data in localStorage
    localStorage.setItem('authToken', token);
    localStorage.setItem('currentUser', JSON.stringify(user));
    this.tokenSubject.next(token);
    this.currentUserSubject.next(user);
  }

  private getStoredToken(): string | null {
    // Retrieve authToken from localStorage
    return localStorage.getItem('authToken');
  }

  private getStoredUser(): User | null {
    try {
      const user = localStorage.getItem('currentUser');
      return user ? JSON.parse(user) : null;
    } catch (e: any) {
      console.error('Error reading from localStorage', e);
      return null;
    }
  }

  private clearAuthData(): void {
    // Clear localStorage data
    localStorage.removeItem('authToken');
    localStorage.removeItem('currentUser');
    this.tokenSubject.next(null);
    this.currentUserSubject.next(null);
  }
}
