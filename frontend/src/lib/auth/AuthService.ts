interface User {
  id: string;
  Name: string;
  email: string;
}

interface LoginResponse {
  token: string;
  user: User;
}

export class AuthService {
  private static instance: AuthService;
  private readonly API_URL = 'http://localhost:8080';
  private readonly TOKEN_KEY = 'token';
  private readonly USER_KEY = 'user';

  private constructor() {}

  public static getInstance(): AuthService {
    if (!AuthService.instance) {
      AuthService.instance = new AuthService();
    }
    return AuthService.instance;
  }

  // Token işlemleri
  public getToken(): string | null {
    if (typeof window === 'undefined') return null;
    return this.getCookie(this.TOKEN_KEY);
  }

  public setToken(token: string): void {
    if (typeof window === 'undefined') return;
    this.setCookie(this.TOKEN_KEY, token, 30); // 30 gün
  }

  public removeToken(): void {
    if (typeof window === 'undefined') return;
    this.removeCookie(this.TOKEN_KEY);
  }

  // Kullanıcı işlemleri
  public getUser(): User | null {
    if (typeof window === 'undefined') return null;
    const userStr = localStorage.getItem(this.USER_KEY);
    return userStr ? JSON.parse(userStr) : null;
  }

  public setUser(user: User): void {
    if (typeof window === 'undefined') return;
    localStorage.setItem(this.USER_KEY, JSON.stringify(user));
  }

  public removeUser(): void {
    if (typeof window === 'undefined') return;
    localStorage.removeItem(this.USER_KEY);
  }

  // Authentication durumu
  public isAuthenticated(): boolean {
    return !!this.getToken();
  }

  // Login işlemi
  public async login(email: string, password: string): Promise<LoginResponse> {
    try {
      const response = await fetch(`${this.API_URL}/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      });

      if (!response.ok) {
        throw new Error('Giriş başarısız');
      }

      const data: LoginResponse = await response.json();

      this.setToken(data.token);
      this.setUser(data.user);

      return data;
    } catch (error) {
      throw new Error('Giriş işlemi başarısız oldu');
    }
  }

  // Logout işlemi
  public logout(): void {
    this.removeToken();
    this.removeUser();
  }

  // API istekleri için yardımcı metod
  public async fetchWithAuth(endpoint: string, options: RequestInit = {}): Promise<any> {
    const token = this.getToken();
    
    const headers = {
      'Accept': 'application/json',
      'Authorization': `Bearer ${token}`,
      ...options.headers,
    };

    const response = await fetch(`${this.API_URL}${endpoint}`, {
      ...options,
      headers,
    });

    if (!response.ok) {
      throw new Error('API isteği başarısız oldu');
    }

    return response.json();
  }

  // Cookie yardımcı metodları
  private setCookie(name: string, value: string, days: number): void {
    const maxAge = days * 24 * 60 * 60;
    document.cookie = `${name}=${value}; path=/; max-age=${maxAge}`;
  }

  private getCookie(name: string): string | null {
    const value = document.cookie
      .split('; ')
      .find(row => row.startsWith(`${name}=`))
      ?.split('=')[1];
    return value || null;
  }

  private removeCookie(name: string): void {
    document.cookie = `${name}=; path=/; expires=Thu, 01 Jan 1970 00:00:01 GMT`;
  }
} 