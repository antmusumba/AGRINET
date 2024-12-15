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

  private currentUserSubject = new BehaviorSubject<User | null>(null);
  currentUser$ = this.currentUserSubject.asObservable();

  private tokenSubject = new BehaviorSubject<string | null>(null);

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
    this.tokenSubject.next(token);
    this.currentUserSubject.next(user);
  }

  private clearAuthData(): void {
    this.tokenSubject.next(null);
    this.currentUserSubject.next(null);
  }
}
