import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface Product {
  id: number;
  product_code: string;
  created_at: string;
  updated_at: string;
}

@Injectable({ providedIn: 'root' })
export class ProductService {
  // relative URL — nginx proxies /api → backend:3000
  private apiUrl = '/api';

  constructor(private http: HttpClient) {}

  getProducts(): Observable<Product[]> {
    return this.http.get<Product[]>(`${this.apiUrl}/products`);
  }

  addProduct(productCode: string): Observable<Product> {
    return this.http.post<Product>(`${this.apiUrl}/products`, { product_code: productCode });
  }

  deleteProduct(id: number): Observable<any> {
    return this.http.delete(`${this.apiUrl}/products/${id}`);
  }
}
