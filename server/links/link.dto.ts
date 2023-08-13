import { IsDate, IsOptional, IsString } from "class-validator";

export class CreateLinkDto {
  @IsString()
  public name: string;

  @IsString() @IsOptional()
  public alias: string;

  @IsString()
  public originalUrl: string;

  @IsDate() @IsOptional()
  public expiryAt: Date;
}
