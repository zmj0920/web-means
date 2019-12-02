import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { AppRoutingModule } from './routes/app-routing.module';
import { AppComponent } from './pages/app/app.component';
import { AppStoreModule } from '@store/store.module';
import { Book1Component } from './component/book1/book1.component';
import { Book2Component } from './component/book2/book2.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { LayoutComponent } from './pages/layout/layout.component';
import { Book3Component } from './component/book3/book3.component';
import { SexPipe } from './utils/sex.pipe';
import { TestDirective } from './utils/test.directive';
import { NgZorroAntdModule, NZ_I18N, zh_CN } from 'ng-zorro-antd';
import { NzFormModule } from 'ng-zorro-antd/form';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { registerLocaleData } from '@angular/common';
import zh from '@angular/common/locales/zh';
import { ProfileEditorComponent } from './component/profile-editor/profile-editor.component';

registerLocaleData(zh);


@NgModule({
  declarations: [AppComponent, Book1Component, Book2Component, LayoutComponent, Book3Component, SexPipe, TestDirective, ProfileEditorComponent],
  // tslint:disable-next-line: max-line-length
  imports: [BrowserModule, AppRoutingModule, AppStoreModule, FormsModule, ReactiveFormsModule, NgZorroAntdModule, HttpClientModule, BrowserAnimationsModule],
  providers: [{ provide: NZ_I18N, useValue: zh_CN }],
  bootstrap: [AppComponent]
})
export class AppModule { }

