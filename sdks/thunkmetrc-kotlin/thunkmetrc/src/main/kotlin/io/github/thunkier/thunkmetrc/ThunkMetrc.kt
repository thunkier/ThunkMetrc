package io.github.thunkier.thunkmetrc

import io.github.thunkier.thunkmetrc.wrapper.MetrcWrapper
import io.github.thunkier.thunkmetrc.wrapper.models.PackagesPackage
import java.time.OffsetDateTime
import java.time.format.DateTimeFormatter

class ThunkMetrc(private val wrapper: MetrcWrapper) {
    private val formatter = DateTimeFormatter.ofPattern("yyyy-MM-dd'T'HH:mm:ssXXX")

    suspend fun activePackagesInventorySync(
        licenseNumber: String,
        lastModifiedStart: OffsetDateTime? = null,
        bufferMinutes: Int? = null
    ): List<PackagesPackage> {
        val allPackages = mutableListOf<PackagesPackage>()
        var pageNumber = 1
        val pageSize = 20

        val startStr = lastModifiedStart?.format(formatter)
        val endStr = if (bufferMinutes != null && bufferMinutes > 0) {
             OffsetDateTime.now().minusMinutes(bufferMinutes.toLong()).format(formatter)
        } else null

        while (true) {
            val result = wrapper.packages.getActivePackages(
                lastModifiedEnd = endStr,
                lastModifiedStart = startStr,
                licenseNumber = licenseNumber,
                pageNumber = pageNumber.toString(),
                pageSize = pageSize.toString()
            )

            val data = result?.Data
            if (data.isNullOrEmpty()) break
            
            allPackages.addAll(data)
            pageNumber++
        }
        return allPackages
    }
}

