

import {HttpParams} from '@angular/common/http';
import {ChangeDetectionStrategy, ChangeDetectorRef, Component, Input} from '@angular/core';
import {Observable} from 'rxjs/Observable';
import {Ingress, IngressList} from 'typings/backendapi';

import {ResourceListBase} from '../../../resources/list';
import {NotificationsService} from '../../../services/global/notifications';
import {EndpointManager, Resource} from '../../../services/resource/endpoint';
import {NamespacedResourceService} from '../../../services/resource/resource';
import {MenuComponent} from '../../list/column/menu/component';
import {ListGroupIdentifier, ListIdentifier} from '../groupids';

@Component({
  selector: 'kd-ingress-list',
  templateUrl: './template.html',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class IngressListComponent extends ResourceListBase<IngressList, Ingress> {
  @Input() endpoint = EndpointManager.resource(Resource.ingress, true).list();

  constructor(
    private readonly ingress_: NamespacedResourceService<IngressList>,
    notifications: NotificationsService,
    cdr: ChangeDetectorRef,
  ) {
    super('ingress', notifications, cdr);
    this.id = ListIdentifier.ingress;
    this.groupId = ListGroupIdentifier.discovery;

    // Register action columns.
    this.registerActionColumn<MenuComponent>('menu', MenuComponent);

    // Register dynamic columns.
    this.registerDynamicColumn('namespace', 'name', this.shouldShowNamespaceColumn_.bind(this));
  }

  getResourceObservable(params?: HttpParams): Observable<IngressList> {
    return this.ingress_.get(this.endpoint, undefined, undefined, params);
  }

  map(ingressList: IngressList): Ingress[] {
    return ingressList.items;
  }

  getDisplayColumns(): string[] {
    return ['name', 'labels', 'endpoints', 'created'];
  }

  private shouldShowNamespaceColumn_(): boolean {
    return this.namespaceService_.areMultipleNamespacesSelected();
  }
}
