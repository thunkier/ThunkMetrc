package com.thunkmetrc

import com.thunkmetrc.wrapper.MetrcWrapper
import java.time.OffsetDateTime
import java.time.format.DateTimeFormatter

class ThunkMetrc(private val wrapper: MetrcWrapper) {
    private val formatter = DateTimeFormatter.ofPattern("yyyy-MM-dd'T'HH:mm:ssXXX")

    suspend fun activePackagesInventorySync(
        licenseNumber: String,
        lastModifiedStart: OffsetDateTime? = null,
        bufferMinutes: Int? = null
    ): List<Any> {
        val allPackages = mutableListOf<Any>()
        var pageNumber = 1
        val pageSize = 20

        val startStr = lastModifiedStart?.format(formatter)
        val endStr = if (bufferMinutes != null && bufferMinutes > 0) {
             OffsetDateTime.now().minusMinutes(bufferMinutes.toLong()).format(formatter)
        } else null

        while (true) {
            // Using named arguments for clarity and safety against parameter ordering
            val result = wrapper.packagesGetActiveV2(
                lastmodifiedend = endStr,
                lastmodifiedstart = startStr,
                licensenumber = licenseNumber,
                pagenumber = pageNumber.toString(),
                pagesize = pageSize.toString()
            )

            if (result == null) break

            if (result is List<*>) {
                if (result.isEmpty()) break
                allPackages.addAll(result.filterNotNull())
                pageNumber++
            } else {
                break
            }
        }
        return allPackages
    }
}
