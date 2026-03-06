import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ProductService, Product } from './product.service';
import { BarcodeComponent } from './barcode.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, FormsModule, BarcodeComponent],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  title = 'IT 06-1';
  productCode = '';
  products: Product[] = [];
  errorMessage = '';
  showConfirm = false;
  pendingDeleteId: number | null = null;
  pendingDeleteCode = '';
  loading = false;

  constructor(private productService: ProductService) {}

  ngOnInit(): void {
    this.loadProducts();
  }

  loadProducts(): void {
    this.productService.getProducts().subscribe({
      next: (data) => this.products = data,
      error: (err) => this.errorMessage = 'Failed to load products'
    });
  }

  onInputChange(): void {
    // Auto-format: insert dashes at positions 4, 9, 14
    let raw = this.productCode.replace(/-/g, '').toUpperCase();
    raw = raw.replace(/[^A-Z0-9]/g, '');
    if (raw.length > 16) raw = raw.substring(0, 16);
    let formatted = '';
    for (let i = 0; i < raw.length; i++) {
      if (i > 0 && i % 4 === 0) formatted += '-';
      formatted += raw[i];
    }
    this.productCode = formatted;
    this.errorMessage = '';
  }

  addProduct(): void {
    if (!this.productCode) {
      this.errorMessage = 'กรุณากรอกรหัสสินค้า';
      return;
    }
    const codeWithoutDash = this.productCode.replace(/-/g, '');
    if (codeWithoutDash.length !== 16) {
      this.errorMessage = 'รหัสสินค้าต้องมี 16 หลัก';
      return;
    }
    this.loading = true;
    this.productService.addProduct(this.productCode).subscribe({
      next: (product) => {
        this.products.push(product);
        this.productCode = '';
        this.errorMessage = '';
        this.loading = false;
      },
      error: (err) => {
        this.errorMessage = err.error?.error || 'เกิดข้อผิดพลาด';
        this.loading = false;
      }
    });
  }

  confirmDelete(product: Product): void {
    this.pendingDeleteId = product.id;
    this.pendingDeleteCode = product.product_code;
    this.showConfirm = true;
  }

  cancelDelete(): void {
    this.showConfirm = false;
    this.pendingDeleteId = null;
    this.pendingDeleteCode = '';
  }

  executeDelete(): void {
    if (this.pendingDeleteId === null) return;
    this.productService.deleteProduct(this.pendingDeleteId).subscribe({
      next: () => {
        this.products = this.products.filter(p => p.id !== this.pendingDeleteId);
        this.showConfirm = false;
        this.pendingDeleteId = null;
      },
      error: (err) => {
        this.errorMessage = 'ลบข้อมูลไม่สำเร็จ';
        this.showConfirm = false;
      }
    });
  }
}
