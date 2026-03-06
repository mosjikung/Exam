import { Component, Input, OnChanges, ElementRef, ViewChild, AfterViewInit } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-barcode',
  standalone: true,
  imports: [CommonModule],
  template: `<canvas #barcodeCanvas></canvas>`
})
export class BarcodeComponent implements OnChanges, AfterViewInit {
  @Input() value = '';
  @ViewChild('barcodeCanvas') canvasRef!: ElementRef<HTMLCanvasElement>;

  // Code 39 encoding: each character = 5 bars + 4 spaces (9 elements)
  // W=wide, n=narrow. Binary: 1=bar, 0=space
  private CODE39_CHARS: { [key: string]: string } = {
    '0': 'nnnWWnWnn', '1': 'WnnWnnnnn', '2': 'nWnWnnnnn',
    '3': 'WWnWnnnnn', '4': 'nnnWWnnnn', '5': 'WnnWWnnnn',
    '6': 'nWnWWnnnn', '7': 'nnnWnWnnn', '8': 'WnnWnWnnn',
    '9': 'nWnWnWnnn', 'A': 'WnnnnnWnn', 'B': 'nWnnnnWnn',
    'C': 'WWnnnnWnn', 'D': 'nnWnnnWnn', 'E': 'WnWnnnWnn',
    'F': 'nWWnnnWnn', 'G': 'nnnWnnWnn', 'H': 'WnnWnnWnn',
    'I': 'nWnWnnWnn', 'J': 'nnWWnnWnn', 'K': 'WnnnWnWnn',
    'L': 'nWnnWnWnn', 'M': 'WWnnWnWnn', 'N': 'nnWnWnWnn',
    'O': 'WnWnWnWnn', 'P': 'nWWnWnWnn', 'Q': 'nnnNnnWnn',
    'R': 'WnnNnnWnn', 'S': 'nWnNnnWnn', 'T': 'nnWNnnWnn',
    'U': 'WnnnnnWWn', 'V': 'nWnnnnWWn', 'W': 'WWnnnnWWn',
    'X': 'nnWnnnWWn', 'Y': 'WnWnnnWWn', 'Z': 'nWWnnnWWn',
    '-': 'nnnWnnWWn', '.': 'WnnWnnWWn', ' ': 'nWnWnnWWn',
    '$': 'nnnnnWWWn', '/': 'nnnWnWWWn', '+': 'nnnnnWWWn',
    '%': 'WnnnWnWWn', '*': 'nnnWnWWWn'
  };

  ngAfterViewInit(): void {
    this.draw();
  }

  ngOnChanges(): void {
    setTimeout(() => this.draw(), 0);
  }

  private draw(): void {
    if (!this.canvasRef?.nativeElement) return;
    const canvas = this.canvasRef.nativeElement;
    const ctx = canvas.getContext('2d');
    if (!ctx) return;

    const narrow = 1.5;
    const wide = 3.5;
    const height = 40;
    const quiet = 10;

    
    const text = '*' + this.value.replace(/-/g, '') + '*';
    const barWidths: number[] = [];
    const gap = narrow; 

    for (let ci = 0; ci < text.length; ci++) {
      const ch = text[ci];
      const pattern = this.CODE39_CHARS[ch] || this.CODE39_CHARS['*'];
      if (!pattern) continue;
      for (let i = 0; i < pattern.length; i++) {
        barWidths.push(pattern[i] === 'W' || pattern[i] === 'N' ? wide : narrow);
      }
      if (ci < text.length - 1) barWidths.push(gap);
    }

    const totalWidth = quiet * 2 + barWidths.reduce((a, b) => a + b, 0);
    canvas.width = Math.ceil(totalWidth);
    canvas.height = height;

    ctx.fillStyle = 'white';
    ctx.fillRect(0, 0, canvas.width, canvas.height);

    let x = quiet;
    let isBar = true; 
    ctx.fillStyle = 'black';
    for (const w of barWidths) {
      if (isBar) {
        ctx.fillRect(Math.round(x), 0, Math.ceil(w), height);
      }
      x += w;
      isBar = !isBar;
    }
  }
}
