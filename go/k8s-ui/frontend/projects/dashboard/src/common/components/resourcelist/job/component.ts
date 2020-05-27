

import {HttpParams} from '@angular/common/http';
import {
  ChangeDetectionStrategy,
  ChangeDetectorRef,
  Component,
  ComponentFactoryResolver,
  Input,
} from '@angular/core';
import {Event, Job, JobList, Metric} from '@api/backendapi';
import {Observable} from 'rxjs/Observable';

import {ResourceListWithStatuses} from '../../../resources/list';
import {NotificationsService} from '../../../services/global/notifications';
import {EndpointManager, Resource} from '../../../services/resource/endpoint';
import {NamespacedResourceService} from '../../../services/resource/resource';
import {MenuComponent} from '../../list/column/menu/component';
import {ListGroupIdentifier, ListIdentifier} from '../groupids';

@Component({
  selector: 'kd-job-list',
  templateUrl: './template.html',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class JobListComponent extends ResourceListWithStatuses<JobList, Job> {
  @Input() title: string;
  @Input() endpoint = EndpointManager.resource(Resource.job, true).list();
  @Input() showMetrics = false;
  cumulativeMetrics: Metric[];

  constructor(
    private readonly job_: NamespacedResourceService<JobList>,
    notifications: NotificationsService,
    resolver: ComponentFactoryResolver,
    cdr: ChangeDetectorRef,
  ) {
    super('job', notifications, cdr, resolver);
    this.id = ListIdentifier.job;
    this.groupId = ListGroupIdentifier.workloads;

    // Register status icon handlers
    this.registerBinding(this.icon.checkCircle, 'kd-success', this.isInSuccessState);
    this.registerBinding(this.icon.timelapse, 'kd-muted', this.isInPendingState);
    this.registerBinding(this.icon.error, 'kd-error', this.isInErrorState);

    // Register action columns.
    this.registerActionColumn<MenuComponent>('menu', MenuComponent);

    // Register dynamic columns.
    this.registerDynamicColumn('namespace', 'name', this.shouldShowNamespaceColumn_.bind(this));
  }

  getResourceObservable(params?: HttpParams): Observable<JobList> {
    return this.job_.get(this.endpoint, undefined, undefined, params);
  }

  map(jobList: JobList): Job[] {
    this.cumulativeMetrics = jobList.cumulativeMetrics;
    return jobList.jobs;
  }

  isInErrorState(resource: Job): boolean {
    return resource.podInfo.warnings.length > 0;
  }

  isInPendingState(resource: Job): boolean {
    return resource.podInfo.warnings.length === 0 && resource.podInfo.pending > 0;
  }

  isInSuccessState(resource: Job): boolean {
    return resource.podInfo.warnings.length === 0 && resource.podInfo.pending === 0;
  }

  getDisplayColumns(): string[] {
    return ['statusicon', 'name', 'labels', 'pods', 'created', 'images'];
  }

  hasErrors(job: Job): boolean {
    return job.podInfo.warnings.length > 0;
  }

  getEvents(job: Job): Event[] {
    return job.podInfo.warnings;
  }

  private shouldShowNamespaceColumn_(): boolean {
    return this.namespaceService_.areMultipleNamespacesSelected();
  }
}
