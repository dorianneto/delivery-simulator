import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { Route, RouteDocument } from './entities/route.entity';

@Injectable()
export class RoutesService {
  constructor(@InjectModel(Route.name) private route: Model<RouteDocument>) {}

  findAll() {
    return this.route.find().exec();
  }
}
