

import {CommonModule} from '@angular/common';
import {NgModule} from '@angular/core';
import {FlexLayoutModule} from '@angular/flex-layout';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {MatButtonModule} from '@angular/material/button';
import {MatButtonToggleModule} from '@angular/material/button-toggle';
import {MatCardModule} from '@angular/material/card';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatChipsModule} from '@angular/material/chips';
import {MatDialogModule} from '@angular/material/dialog';
import {MatDividerModule} from '@angular/material/divider';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatIconModule} from '@angular/material/icon';
import {MatInputModule} from '@angular/material/input';
import {MatMenuModule} from '@angular/material/menu';
import {MatPaginatorModule} from '@angular/material/paginator';
import {MatProgressBarModule} from '@angular/material/progress-bar';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {MatRadioModule} from '@angular/material/radio';
import {MatSelectModule} from '@angular/material/select';
import {MatSidenavModule} from '@angular/material/sidenav';
import {MatSlideToggleModule} from '@angular/material/slide-toggle';
import {MatSliderModule} from '@angular/material/slider';
import {MatSnackBarModule} from '@angular/material/snack-bar';
import {MatSortModule} from '@angular/material/sort';
import {MatTableModule} from '@angular/material/table';
import {MatTabsModule} from '@angular/material/tabs';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MAT_TOOLTIP_DEFAULT_OPTIONS, MatTooltipModule} from '@angular/material/tooltip';
import {RouterModule} from '@angular/router';
import {FilterPipeModule} from 'ngx-filter-pipe';

import {PipesModule} from './common/pipes/module';
import {KD_TOOLTIP_DEFAULT_OPTIONS} from './index.config';
import {NgxChartsModule} from '@swimlane/ngx-charts';

const SHARED_DEPENDENCIES = [
  // Angular imports
  CommonModule,
  FormsModule,
  ReactiveFormsModule,

  // Material imports
  MatButtonModule,
  MatCardModule,
  MatCheckboxModule,
  MatDividerModule,
  MatGridListModule,
  MatFormFieldModule,
  MatIconModule,
  MatInputModule,
  MatProgressSpinnerModule,
  MatProgressBarModule,
  MatRadioModule,
  MatSidenavModule,
  MatToolbarModule,
  MatTooltipModule,
  MatSliderModule,
  MatDialogModule,
  MatSlideToggleModule,
  MatChipsModule,
  MatTableModule,
  MatSortModule,
  MatPaginatorModule,
  MatTabsModule,
  MatMenuModule,
  MatSelectModule,
  MatButtonToggleModule,
  MatSnackBarModule,

  // Other 3rd party modules
  FlexLayoutModule,
  RouterModule,
  FilterPipeModule,
  NgxChartsModule,

  // Custom application modules
  PipesModule,
];

@NgModule({
  imports: SHARED_DEPENDENCIES,
  exports: SHARED_DEPENDENCIES,
  providers: [
    {
      provide: MAT_TOOLTIP_DEFAULT_OPTIONS,
      useValue: KD_TOOLTIP_DEFAULT_OPTIONS,
    },
  ],
})
export class SharedModule {}
