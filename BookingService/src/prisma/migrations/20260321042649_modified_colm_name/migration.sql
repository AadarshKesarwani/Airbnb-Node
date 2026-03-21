/*
  Warnings:

  - You are about to drop the column `key` on the `idempotencykey` table. All the data in the column will be lost.
  - A unique constraint covering the columns `[IdemKey]` on the table `IdempotencyKey` will be added. If there are existing duplicate values, this will fail.
  - Added the required column `IdemKey` to the `IdempotencyKey` table without a default value. This is not possible if the table is not empty.
  - Made the column `bookingId` on table `idempotencykey` required. This step will fail if there are existing NULL values in that column.

*/
-- DropForeignKey
ALTER TABLE `idempotencykey` DROP FOREIGN KEY `IdempotencyKey_bookingId_fkey`;

-- DropIndex
DROP INDEX `IdempotencyKey_key_key` ON `idempotencykey`;

-- AlterTable
ALTER TABLE `idempotencykey` DROP COLUMN `key`,
    ADD COLUMN `IdemKey` VARCHAR(191) NOT NULL,
    MODIFY `bookingId` INTEGER NOT NULL;

-- CreateIndex
CREATE UNIQUE INDEX `IdempotencyKey_IdemKey_key` ON `IdempotencyKey`(`IdemKey`);

-- AddForeignKey
ALTER TABLE `IdempotencyKey` ADD CONSTRAINT `IdempotencyKey_bookingId_fkey` FOREIGN KEY (`bookingId`) REFERENCES `Booking`(`id`) ON DELETE RESTRICT ON UPDATE CASCADE;
