export * from './auth.service';
import { AuthService } from './auth.service';
export * from './products.service';
import { ProductsService } from './products.service';
export * from './sales.service';
import { SalesService } from './sales.service';
export * from './settings.service';
import { SettingsService } from './settings.service';
export const APIS = [AuthService, ProductsService, SalesService, SettingsService];
