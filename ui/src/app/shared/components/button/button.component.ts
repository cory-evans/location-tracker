import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import type { IconDefinition } from '@fortawesome/fontawesome-common-types';

const VARIANTS = {
  primary:
    'bg-primary-600 text-white hover:bg-primary-200 hover:text-primary-700',
  inverse: 'bg-white text-primary-600 hover:bg-primary-100',
  danger: 'bg-danger-600 text-white hover:bg-danger-200 hover:text-danger-700',
};

const SIZES = {
  sm: 'py-2 px-4 text-sm rounded-sm',
  md: 'py-2 px-6 text-md rounded-lg',
  lg: 'py-3 px-8 text-lg rounded-lg',
};

@Component({
  selector: 'app-button',
  templateUrl: './button.component.html',
})
export class ButtonComponent implements OnInit {
  @Input() variant: keyof typeof VARIANTS = 'primary';
  @Input() size: keyof typeof SIZES = 'md';
  @Input() text?: string;
  @Input() icon?: IconDefinition;

  @Input() type: 'button' | 'menu' | 'submit' | 'reset' = 'button';

  @Output() click = new EventEmitter<MouseEvent>();

  constructor() {}

  ngOnInit(): void {}

  get btnClass() {
    return (
      'transition-colors border ' +
      VARIANTS[this.variant] +
      ' ' +
      SIZES[this.size]
    );
  }
}
