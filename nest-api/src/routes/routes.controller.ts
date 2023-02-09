import { Controller, Get, Inject } from '@nestjs/common';
import { ClientKafka, MessagePattern, Payload } from '@nestjs/microservices';
import { RoutesService } from './routes.service';
import { RoutesGateway } from './routes.gateway';

@Controller('routes')
export class RoutesController {
  constructor(
    private readonly routesService: RoutesService,
    @Inject('KAFKA_SERVICE')
    private clientKafka: ClientKafka,
    private routeGateway: RoutesGateway,
  ) {}

  @Get()
  findAll() {
    return this.routesService.findAll();
  }

  @MessagePattern('route.new-position')
  consumeNewPosition(
    @Payload()
    message: {
      value: {
        routeId: string;
        clientId: string;
        position: [number, number];
        finished: boolean;
      };
    },
  ) {
    this.routeGateway.sendPosition(message.value);
  }
}
