

#include <alpm.h>
#include <stdio.h>

alpm_errno_t handleError(alpm_errno_t err) {
    if (err != 0) {
        fprintf(stderr, "Error %s (%d)\n", alpm_strerror(err), err);
    }
}

int main() {
    alpm_errno_t err;
    alpm_handle_t *handle = alpm_initialize("/", "/var/lib/pacman", &err);
    if (err != 0) {
        return handleError(err);
    }

    alpm_db_t* database = alpm_get_localdb(handle);
    if (!database) {
        err = alpm_errno(handle);
        if (err != 0) {
            alpm_release(handle);
            return handleError(err);
        }
    }

    alpm_list_t* packageList = alpm_db_get_pkgcache(database);
    if (!packageList) {
        err = alpm_errno(handle);
        if (err != 0) {
            alpm_release(handle);
            return handleError(err);
        }
    }

    while (packageList) {
        alpm_pkg_t* package = (alpm_pkg_t*)(packageList->data);
        fprintf(stdout, "%s\n", alpm_pkg_get_name(package));
        packageList = packageList->next;
    }

    return alpm_release(handle);
}
