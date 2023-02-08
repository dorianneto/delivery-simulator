import { Controller, Get, Inject, OnModuleInit, Param } from '@nestjs/common';
import { ClientKafka } from '@nestjs/microservices';
import { Producer } from '@nestjs/microservices/external/kafka.interface';
import { RoutesService } from './routes.service';

@Controller('routes')
export class RoutesController implements OnModuleInit {
  private kafkaProducer: Producer;

  constructor(
    private readonly routesService: RoutesService,
    @Inject('KAFKA_SERVICE') private readonly clientKafka: ClientKafka,
  ) {}

  @Get()
  findAll() {
    return this.routesService.findAll();
  }

  async onModuleInit() {
    this.kafkaProducer = await this.clientKafka.connect();
  }

  @Get(':id/start')
  start(@Param('id') id: string) {
    console.log('sending to kafka...');
    this.kafkaProducer.send({
      topic: 'route.new-direction',
      messages: [
        {
          key: 'route.new-direction',
          value: JSON.stringify({ routeId: id, clientId: '' }),
        },
      ],
    });
  }
}
